using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace AvioApp.Model
{
    [BsonIgnoreExtraElements]
    public class Flight
    {
        [BsonId]
        [BsonRepresentation(BsonType.ObjectId)]
        public string Id { get; set; }
        [BsonElement("date")]
        public DateTime Date { get; set; }
        [BsonElement("duration")]
        public int Duration { get; set; }
        [BsonElement("start")]
        public string Start { get; set; } = string.Empty;
        [BsonElement("destination")]
        public string Destination { get; set; } = string.Empty;
        [BsonElement("price")]
        public float Price { get; set; }
        [BsonElement("tickets")]
        public int Tickets { get; set; }
    }
}
