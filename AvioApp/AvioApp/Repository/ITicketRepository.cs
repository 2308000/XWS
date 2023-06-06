using AvioApp.Model;

namespace AvioApp.Repository
{
    public interface ITicketRepository
    {
        public IQueryable<Ticket> GetAll();
        public Ticket? Create(Ticket entity);
        public Ticket Update(Ticket entity);
        public void Delete(Ticket entity);
    }
}
