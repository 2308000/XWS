using AvioApp.Model;
using MongoDB.Driver;

namespace AvioApp.Repository
{
    public class UserRepository : IUserRepository
    {
        private readonly IMongoCollection<User> _users;
        public UserRepository(IMongoClient mongoClient)
        {
            var db = mongoClient.GetDatabase("XWS_DB");
            _users = db.GetCollection<User>("users");
        }
        public User Create(User entity)
        {
            _users.InsertOne(entity);
            return entity;
        }

        public void Delete(string id)
        {
            _users.DeleteOne(u => u.Id == id);
        }

        public User Get(string id)
        {
            return _users.Find(u => u.Id == id).FirstOrDefault();
        }

        public IEnumerable<User> GetAll()
        {
            return _users.Find(u => true).ToList();
        }

        public void Update(string id, User entity)
        {
            _users.ReplaceOne(u => u.Id == id, entity);
        }
    }
}
