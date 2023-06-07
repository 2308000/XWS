using AvioApp.Model;

namespace AvioApp.Repository
{
    public interface ITicketRepository
    {
        public IEnumerable<Ticket> GetAll();
        public Ticket Get(string id);
        public Ticket Create(Ticket entity);
        public void Update(string id, Ticket entity);
        public void Delete(string id);
    }
}
