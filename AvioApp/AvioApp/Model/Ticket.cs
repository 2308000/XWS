using System.ComponentModel.DataAnnotations.Schema;

namespace AvioApp.Model
{
    public class Ticket
    {
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        public long Id { get; set; }
        public User User { get; set; }
        public Flight Flight { get; set; }
    }
}
