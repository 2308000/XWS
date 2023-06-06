namespace AvioApp.Repository
{
    public interface IUnitOfWork
    {
        IUserRepository UserRepository { get; }
        IFlightRepository FlightRepository { get; }
        ITicketRepository TicketRepository { get; }
        void SaveChanges();
    }
}
