using AvioApp.Model;
using Microsoft.EntityFrameworkCore;

namespace AvioApp.Repository
{
    public class UserRepository : IUserRepository
    {
        private readonly DbContext _dbContext;

        public UserRepository(DbContext dbContext)
        {
            _dbContext = dbContext;
        }
        public User? Create(User entity)
        {
            _dbContext.Set<User>().Add(entity);
            return entity;
        }

        public void Delete(User entity)
        {
            _dbContext.Set<User>().Remove(entity);
        }

        public IQueryable<User> GetAll()
        {
            return _dbContext.Set<User>();
        }

        public User Update(User entity)
        {
            _dbContext.Set<User>().Update(entity);
            return entity;
        }
    }
}
