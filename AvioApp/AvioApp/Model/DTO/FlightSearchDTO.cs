namespace AvioApp.Model.DTO
{
    public class FlightSearchDTO
    {
        public DateTime Date { get; set; }
        public string Start { get; set; } = string.Empty;
        public string Destination { get; set; } = string.Empty;
        public int RequiredTickets { get; set; }
    }
}
