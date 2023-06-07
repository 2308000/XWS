using AvioApp.Model;
using MongoDB.Driver;

namespace AvioApp.Repository
{
    public class FlightRepository : IFlightRepository
    {
        private readonly IMongoCollection<Flight> _flights;
        public FlightRepository(IMongoClient mongoClient)
        {
            var db = mongoClient.GetDatabase("XWS_DB");
            _flights = db.GetCollection<Flight>("flights");
        }
        public Flight Create(Flight entity)
        {
            _flights.InsertOne(entity);
            return entity;
        }

        public void Delete(string id)
        {
            _flights.DeleteOne(f => f.Id == id);
        }

        public Flight Get(string id)
        {
            return _flights.Find(f => f.Id == id).FirstOrDefault();
        }

        public IEnumerable<Flight> GetAll()
        {
            return _flights.Find(u => true).ToList();
        }

        public void Update(string id, Flight entity)
        {
            _flights.ReplaceOne(f => f.Id == id, entity);
        }
    }
}
