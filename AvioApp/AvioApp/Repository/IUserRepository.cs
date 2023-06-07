using AvioApp.Model;

namespace AvioApp.Repository
{
    public interface IUserRepository
    {
        public IEnumerable<User> GetAll();
        public User Get(string id);
        public User Create(User entity);
        public void Update(string id, User entity);
        public void Delete(string id);
    }
}
