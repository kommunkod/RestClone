using Kommunkod.Restclone;

namespace Kommunkod.Restclone.Test;
public class Program
{
    public static void Main(string[] args)
    {
        var filelist = Operations.DirList(new Remote()
        {
            Name = "blabla",
            Type = "azureblob",
            Parameters = new
            {
            },
        }, "", true);

        foreach (var file in filelist.Files)
        {
            Console.WriteLine(file.Name);
        }

        // Console.WriteLine((new AzureRemote()).GetRemoteType());


        // var request = new Interop.RequestFormat
        // {
        //     Method = "POST",
        //     Url = "http://localhost/api/v1/dir/list",
        //     Headers = new Dictionary<string, List<string>>
        //     {
        //         { "Content-Type", new List<string> { "application/json" } }
        //     },
        // };
        //
        // request.SetBody(new
        // {
        //     remote = new
        //     {
        //         name = "myremote",
        //         type = "azureblob",
        //         parameters = new
        //         {
        //         }
        //     },
        //     path = "",
        //     recurse = false,
        // });
        //
        // var response = request.Invoke<Dictionary<string, object>>();
        //
        // Console.WriteLine($"Status Code: {response.StatusCode}");
        // Console.WriteLine($"Status Info: {response.StatusInfo}");
        // Console.WriteLine("Headers:");
        // foreach (var header in response.Headers)
        // {
        //     Console.WriteLine($"  {header.Key}: {string.Join(", ", header.Value)}");
        // }
        //
        // var isError = response.IsError();
        // if (isError != null)
        // {
        //     throw isError;
        // }
        //
        // var body = response.GetBodyObject();
        // Console.WriteLine("Body:");
        // foreach (var kvp in body)
        // {
        //     Console.WriteLine($"  {kvp.Key}: {kvp.Value}");
        // }
    }
}
