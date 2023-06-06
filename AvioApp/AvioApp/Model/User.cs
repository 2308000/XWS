using System.ComponentModel.DataAnnotations.Schema;

namespace AvioApp.Model
{
    public class User
    {
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        public long Id { get; set; }
        public string Email { get; set; } = string.Empty;
        public string Password { get; set; } = string.Empty;
        public byte[] Salt { get; set; } = new byte[0];
        public string FirstName { get; set; } = string.Empty;
        public string LastName { get; set; } = string.Empty;
        public string Role { get; set; } = string.Empty;
        public byte[] Code { get; set; } = new byte[0];
    }
}
