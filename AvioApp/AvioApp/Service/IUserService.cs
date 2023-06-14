using AvioApp.Model.DTO;

namespace AvioApp.Service
{
    public interface IUserService
    {
        public JWTDTO Authenticate(CredentialsDTO credentials);
        public string Register(NewUserDTO user);
        public void UpdateCode(string code, string email);
    }
}
