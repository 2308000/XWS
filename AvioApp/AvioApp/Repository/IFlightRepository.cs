using AvioApp.Model;

namespace AvioApp.Repository
{
    public interface IFlightRepository
    {
        public IEnumerable<Flight> GetAll();
        public Flight Get(string id);
        public Flight Create(Flight entity);
        public void Update(string id, Flight entity);
        public void Delete(string id);
    }
}
