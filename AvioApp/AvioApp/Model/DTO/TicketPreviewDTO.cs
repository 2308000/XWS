namespace AvioApp.Model.DTO
{
    public class TicketPreviewDTO
    {
        public string TicketId { get; set; }
        public string FlightId { get; set; }
        public DateTime Date { get; set; }
        public string Start { get; set; } = string.Empty;
        public string Destination { get; set; } = string.Empty;
        public float Price { get; set; }
    }
}
