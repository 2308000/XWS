namespace AvioApp.Model.DTO
{
    public class FlightUserPreviewDTO
    {
        public string Id { get; set; }
        public DateTime Date { get; set; }
        public int Duration { get; set; }
        public string Start { get; set; } = string.Empty;
        public string Destination { get; set; } = string.Empty;
        public float Price { get; set; }
        public float TotalPrice { get; set; }
    }
}
