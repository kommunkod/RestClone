using System.Runtime.InteropServices;
using System.Text.Json;
using System.Text.Json.Serialization;
using Newtonsoft.Json;

namespace Kommunkod.Restclone;

public class Go
{
    /// <summary>
    /// GoString struct for representing Go strings in C#.
    /// </summary>
    public struct GoString
    {
        public IntPtr p;
        public Int64 n;
    }

    /// <summary>
    /// Operation function for invoking the Go function from C#.
    /// </summary>
    [DllImport("./restclone.dll", CharSet = CharSet.Unicode, CallingConvention = CallingConvention.StdCall, SetLastError = true)]
    [return: MarshalAs(UnmanagedType.LPStr)]
    public static extern string Operation(GoString s);

    public static string Invoke(string requestData)
    {
        GoString gs = new()
        {
            p = Marshal.StringToHGlobalAnsi(requestData),
            n = requestData.Length
        };

        var responseData = Go.Operation(gs);
        return responseData;
    }
}

[Serializable]
public class InvalidConfigurationException : Exception
{
    public InvalidConfigurationException ()
    {}

    public InvalidConfigurationException (string message) 
        : base(message)
    {}

    public InvalidConfigurationException (string message, Exception innerException)
        : base (message, innerException)
    {}    
}


/// <summary>
/// Interop class for handling communication between C# and Go.
/// </summary>
public class Interop
{
    /// <summary>
    /// RequestFormat struct for representing HTTP request data.
    /// </summary>
    public struct RequestFormat
    {
        public string Method;
        public string Url;
        public Dictionary<string, List<string>> Headers;
        public string Body;

        public void SetBody(object body)
        {
            this.Body = Newtonsoft.Json.JsonConvert.SerializeObject(body, new JsonSerializerSettings { 
                NullValueHandling = NullValueHandling.Ignore
            });
        }

        public string Serialize()
        {

            var data = Newtonsoft.Json.JsonConvert.SerializeObject(new
            {
                method = this.Method,
                url = this.Url,
                headers = this.Headers,
                body = this.Body
            }, new JsonSerializerSettings { 
                NullValueHandling = NullValueHandling.Ignore
            });

            var b64data = Convert.ToBase64String(System.Text.Encoding.UTF8.GetBytes(data));
            return b64data;
        }

        /// <summary>
        /// Invoke the HTTP request and get the response.
        /// </summary>
        /// <typeparam name="T">The type of response</typeparam>
        /// <returns>Resposne of type T</returns>
        public ResponseFormat<T> Invoke<T>()
        {
            var requestData = this.Serialize();

            var responseData = Go.Invoke(requestData);

            var response = ResponseFormat<T>.Deserialize(responseData);
            return response;
        }
    }

    /// <summary>
    /// ResponseFormat struct for representing HTTP response data.
    /// </summary>
    /// <typeparam name="T"></typeparam>
    public struct ResponseFormat<T>
    {
        /// <summary>
        /// HTTP status code of the response.
        /// </summary>
        [JsonPropertyName("StatusCode")]
        public int StatusCode { get; set; }

        /// <summary>
        /// HTTP status information of the response.
        /// </summary>
        [JsonPropertyName("StatusInfo")]
        public string StatusInfo { get; set; }

        /// <summary>
        /// HTTP headers of the response.
        /// </summary>
        [JsonPropertyName("Headers")]
        public Dictionary<string, List<string>> Headers { get; set; }

        /// <summary>
        /// HTTP body of the response, encoded in base64.
        /// </summary>
        [JsonPropertyName("Body")]
        public string Body { get; set; }

        private string jsonResponse { get; set; }

        /// <summary>
        /// Deserialize a base64-encoded JSON string into a ResponseFormat object.
        /// </summary>
        /// <param name="data">The base64-encoded JSON string.</param>
        /// <returns>A ResponseFormat object</returns>
        public static ResponseFormat<T> Deserialize(string data)
        {
            var decodedBytes = Convert.FromBase64String(data);
            
            Console.WriteLine(System.Text.Encoding.UTF8.GetString(decodedBytes));

            var reader = new Utf8JsonReader(decodedBytes);

            var obj = System.Text.Json.JsonSerializer.Deserialize<ResponseFormat<T>>(ref reader);
            return obj;
        }

        /// <summary>
        /// Get the body of the response as a UTF-8 string.
        /// </summary>
        /// <returns>The decoded body of the response.</returns>
        public string GetBody()
        {
            var decodedBytes = Convert.FromBase64String(this.Body);
            return System.Text.Encoding.UTF8.GetString(decodedBytes);
        }

        public void DecodeBody()
        {
            if (this.jsonResponse == null)
            {
                var decodedBytes = Convert.FromBase64String(this.Body);
                this.jsonResponse = System.Text.Encoding.UTF8.GetString(decodedBytes);    
            }
        }

        public Exception? IsError()
        {

            this.DecodeBody();
            return this.StatusCode > 299 ? new InvalidConfigurationException(this.jsonResponse) : null;
        }

        /// <summary>
        /// Decode the body of the response into an object of type T.
        /// </summary>
        /// <returns>An object of type T</returns>
        public T? GetBodyObject()
        {
            this.DecodeBody();
            return System.Text.Json.JsonSerializer.Deserialize<T>(this.jsonResponse);
        }
    }
}
