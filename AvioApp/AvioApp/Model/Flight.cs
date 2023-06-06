using System.ComponentModel.DataAnnotations.Schema;

namespace AvioApp.Model
{
    public class Flight
    {
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        public long Id { get; set; }
        public DateTime Date { get; set; }
        public int Duration { get; set; }
        public string Start { get; set; } = string.Empty;
        public string Destination { get; set; } = string.Empty;
        public float Price { get; set; }
        public int Tickets { get; set; }
    }
}
