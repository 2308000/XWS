using AvioApp.Model.DTO;

namespace AvioApp.Service
{
    public interface IUserService
    {
        public string Authenticate(CredentialsDTO credentials);
        public long Register(NewUserDTO user);
        public void UpdateCode(string code, string email);
    }
}
