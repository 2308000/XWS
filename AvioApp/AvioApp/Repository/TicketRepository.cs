using AvioApp.Model;
using MongoDB.Driver;

namespace AvioApp.Repository
{
    public class TicketRepository : ITicketRepository
    {
        private readonly IMongoCollection<Ticket> _tickets;
        public TicketRepository(IMongoClient mongoClient)
        {
            var db = mongoClient.GetDatabase("XWS_DB");
            _tickets = db.GetCollection<Ticket>("tickets");
        }
        public Ticket Create(Ticket entity)
        {
            _tickets.InsertOne(entity);
            return entity;
        }

        public void Delete(string id)
        {
            _tickets.DeleteOne(t => t.Id == id);
        }

        public Ticket Get(string id)
        {
            return _tickets.Find(t => t.Id == id).FirstOrDefault();
        }

        public IEnumerable<Ticket> GetAll()
        {
            return _tickets.Find(t => true).ToList();
        }

        public void Update(string id, Ticket entity)
        {
            _tickets.ReplaceOne(t => t.Id == id, entity);
        }
    }
}
