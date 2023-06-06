using AvioApp.Model;

namespace AvioApp.Repository
{
    public interface IFlightRepository
    {
        public IQueryable<Flight> GetAll();
        public Flight? Create(Flight entity);
        public Flight Update(Flight entity);
        public void Delete(Flight entity);
    }
}
