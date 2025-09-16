using System.Runtime.InteropServices;
using System.Text;
using System.Net;
using System.Net.Sockets;
using System.Threading.Tasks;
using System.IO;
using System;
using System.Collections.Generic;
using System.Linq;


// See https://aka.ms/new-console-template for more information


// Console.WriteLine("Hello, World!");
// StartSocket("/tmp/restclone.sock");

class Program
{
    [DllImport("./restclone", EntryPoint = "StartSocket")]
    extern static int StartSocket(string path);


    static async Task Main(string[] args)
    {
        var task = Task.Run(()  => {
            StartSocket("/tmp/restclone.sock");
            
            // Remove socket if exists
            if (File.Exists("/tmp/restclone.sock"))
            {
                File.Delete("/tmp/restclone.sock");
            }
        });
        Console.WriteLine("Waiting for socket to be ready...");
        await Task.Delay(5500);
        Console.WriteLine("Socket should be ready now.");  


        // Console.WriteLine("Hello, World!");
        // StartSocket("/tmp/restclone.sock");

        // await Task.Delay(-1);

        // Get /docs/ on socket addr
        using (var client = new Socket(AddressFamily.Unix, SocketType.Stream, 0))
        {
            client.Connect(new UnixDomainSocketEndPoint("/tmp/restclone.sock"));
            // client.Connect("/tmp/restclone.sock");
            using (var stream = new NetworkStream(client))
            using (var reader = new StreamReader(stream))
            using (var writer = new StreamWriter(stream) { AutoFlush = true })
            {
                var request = "GET /docs/ HTTP/1.1\r\nHost: localhost\r\n\r\n";
                await writer.WriteAsync(request);
                var response = await reader.ReadToEndAsync();
                Console.WriteLine(response);
            }
        }
    }
}
