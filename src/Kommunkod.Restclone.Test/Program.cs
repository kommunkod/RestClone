using Kommunkod.Restclone;

internal class Program
{
    private static void Main(string[] args)
    {
        var request = new Interop.RequestFormat
        {
            Method = "POST",
            Url = "http://localhost/api/v1/dir/list",
            Headers = new Dictionary<string, List<string>>
            {
                { "Content-Type", new List<string> { "application/json" } }
            },
        };

        request.SetBody(new
        {
            remote = new
            {
                name = "myremote",
                type = "local",
                parameters = new
                {
                    nounc = true
                }
            },
            path = "",
            recurse = false,
        });

        var response = request.Invoke<Dictionary<string, object>>();

        Console.WriteLine($"Status Code: {response.StatusCode}");
        Console.WriteLine($"Status Info: {response.StatusInfo}");
        Console.WriteLine("Headers:");
        foreach (var header in response.Headers)
        {
            Console.WriteLine($"  {header.Key}: {string.Join(", ", header.Value)}");
        }

        var body = response.DecodeBody();
        Console.WriteLine("Body:");
        foreach (var kvp in body)
        {
            Console.WriteLine($"  {kvp.Key}: {kvp.Value}");
        }
    }
}

// using System.Runtime.InteropServices;
// using System.Text.Json;
// using System.Text.Json.Serialization;


// // See https://aka.ms/new-console-template for more information


// // Console.WriteLine("Hello, World!");
// // StartSocket("/tmp/restclone.sock");

// class Go {
//     public struct GoString
//     {
//         public IntPtr p;
//         public Int64 n;
//     }

//     public struct RequestFormat
//     {
//         public string Method;
//         public string Url;
//         public Dictionary<string, List<string>> Headers;
//         public string Body;

//         public string Serialize()
//         {
//             var data = JsonSerializer.Serialize(new
//             {
//                 method = this.Method,
//                 url = this.Url,
//                 headers = this.Headers,
//                 body = this.Body
//             });

//             var b64data = Convert.ToBase64String(System.Text.Encoding.UTF8.GetBytes(data));
//             return b64data;
//         }
//     }

//     public struct ResponseFormat
//     {
//         [JsonPropertyName("StatusCode")]
//         public int StatusCode { get; set;  }
        
//         [JsonPropertyName("StatusInfo")]
//         public string StatusInfo { get; set;  }
        
//         [JsonPropertyName("Headers")]
//         public Dictionary<string, List<string>> Headers { get; set;  }
        
//         [JsonPropertyName("Body")]
//         public string Body { get; set;  }

//         public static ResponseFormat Deserialize(string data)
//         {
//             var decodedBytes = Convert.FromBase64String(data);
//             var obj = JsonSerializer.Deserialize<ResponseFormat>(System.Text.Encoding.UTF8.GetString(decodedBytes));
//             return obj;
//         }

//         public string GetBody()
//         {
//             var decodedBytes = Convert.FromBase64String(this.Body);
//             return System.Text.Encoding.UTF8.GetString(decodedBytes);
//         }
//     }

//     static class Func
//     {
//         [DllImport("../Kommunkod.Restclone/restclone.dll", CharSet = CharSet.Unicode, CallingConvention = CallingConvention.StdCall, SetLastError = true)]
//         [return: MarshalAs(UnmanagedType.LPStr)] 
//         public static extern string Operation(GoString s);
//     }

//     public static ResponseFormat Exec(RequestFormat request) {
//         var input = request.Serialize();
//         Console.WriteLine(input);
//         GoString s = new GoString
//         {  
//             p = Marshal.StringToHGlobalAnsi(input),
//             n = input.Length
//         };

//         var retstr = Func.Operation(s);

//         var response = ResponseFormat.Deserialize(retstr);

//         Marshal.FreeHGlobal(s.p);

//         return response;
//     }
// }

// class Program
// {
//     static void Main(string[] args)
//     {
//         var req = new Go.RequestFormat
//         {
//             Method = "POST",
//             Url = "http://localhost/api/v1/dir/list",
//             Headers = new Dictionary<string, List<string>>
//             {
//                 { "Content-Type", new List<string> { "application/json" } }
//             },
//         };

//         var Body = new {
//             remote = new {
//                 name = "myremote",
//                 type = "local",
//                 parameters = new
//                 {
//                     nounc = true
//                 }
//             },
//             path = "",
//             recurse = false,
//         };

//         var serialized = JsonSerializer.Serialize(Body);
//         Console.WriteLine(serialized);
//         req.Body = serialized;

//         var response = Go.Exec(req);

//         Console.WriteLine("Response: ", response);

//         Console.WriteLine($"Status: {response.StatusCode} {response.StatusInfo}");
//         Console.WriteLine("Headers:");
//         foreach (var header in response.Headers)
//         {
//             Console.WriteLine($"{header.Key}: {string.Join(", ", header.Value)}");
//         }

//         Console.WriteLine("Body:");
//         Console.WriteLine(response.GetBody());
//     }
// }
