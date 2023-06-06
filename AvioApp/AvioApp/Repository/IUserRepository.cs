using AvioApp.Model;

namespace AvioApp.Repository
{
    public interface IUserRepository
    {
        public IQueryable<User> GetAll();
        public User? Create(User entity);
        public User Update(User entity);
        public void Delete(User entity);
    }
}
