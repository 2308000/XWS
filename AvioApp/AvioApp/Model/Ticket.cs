using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace AvioApp.Model
{
    [BsonIgnoreExtraElements]
    public class Ticket
    {
        [BsonId]
        [BsonRepresentation(BsonType.ObjectId)]
        public string Id { get; set; }
        [BsonElement("userid")]
        [BsonRepresentation(BsonType.ObjectId)]
        public string UserId { get; set; }
        [BsonElement("flightid")]
        [BsonRepresentation(BsonType.ObjectId)]
        public string FlightId { get; set; }
    }
}
