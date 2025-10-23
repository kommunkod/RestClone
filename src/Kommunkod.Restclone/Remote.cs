using System.Text.Json.Serialization;

namespace Kommunkod.Restclone;

public class Remote
{
    [JsonPropertyName("name")]
    public string Name { get; set; }
    
    [JsonPropertyName("type")]
    public string Type { get; set; }

    [JsonPropertyName("parameters")] public object Parameters { get; set; } = new { };
    
    [JsonPropertyName("options")]
    public object Options { get; set; }
}