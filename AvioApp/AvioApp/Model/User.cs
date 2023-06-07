using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace AvioApp.Model
{
    [BsonIgnoreExtraElements]
    public class User
    {
        [BsonId]
        [BsonRepresentation(BsonType.ObjectId)]
        public string Id { get; set; }
        [BsonElement("email")]
        public string Email { get; set; } = string.Empty;
        [BsonElement("password")]
        public string Password { get; set; } = string.Empty;
        [BsonElement("salt")]
        public byte[] Salt { get; set; } = new byte[0];
        [BsonElement("firstname")]
        public string FirstName { get; set; } = string.Empty;
        [BsonElement("lastname")]
        public string LastName { get; set; } = string.Empty;
        [BsonElement("role")]
        public string Role { get; set; } = string.Empty;
        [BsonElement("code")]
        public byte[] Code { get; set; } = new byte[0];
    }
}
