using AvioApp.Model;
using AvioApp.Model.DTO;
using AvioApp.Repository;
using AvioApp.SupportClasses;
using AvioApp.SupportClasses.GEH.CustomExceptions;
using System.Text;

namespace AvioApp.Service
{
    public class UserService : IUserService
    {
        private readonly IUserRepository _userRepository;
        private readonly IJWTGenerator _jwtGenerator;

        public UserService(IUserRepository userRepository, IJWTGenerator jwtGenerator)
        {
            _userRepository = userRepository;
            _jwtGenerator = jwtGenerator;
        }

        public JWTDTO Authenticate(CredentialsDTO credentials)
        {
            User? user = _userRepository.GetAll().FirstOrDefault(u => u.Email == credentials.Email);
            if (user == null)
            {
                throw new NotFoundException($"User with email {credentials.Email} does not exists!");
            }
            if (PasswordHasher.VerifyPassword(credentials.Password, user.Password, user.Salt))
            {
                return new JWTDTO { Token = _jwtGenerator.GenerateToken(user, credentials.IsLoginPermanent) };
            }
            throw new BadCredentialsException($"Incorrect password for user with email {credentials.Email}!");
        }

        public string Register(NewUserDTO user)
        {
            if (_userRepository.GetAll().Where(u => u.Email == user.Email).Any())
            {
                throw new DuplicateItemException($"Email {user.Email} is already taken!");
            }
            byte[] salt;
            var newUser = new User
            {
                Email = user.Email,
                Password = PasswordHasher.HashPassword(user.Password, out salt),
                Salt = salt,
                Role = "USER",
                FirstName = user.FirstName,
                LastName = user.LastName
            };
            newUser = _userRepository.Create(newUser);
            return newUser.Id;
        }

        public void UpdateCode(string code, string email)
        {
            var user = _userRepository.GetAll().FirstOrDefault(u => u.Email == email);
            if (user == null)
            {
                throw new NotFoundException($"User with email: {email} does not exists!");
            }
            user.Code = Encoding.ASCII.GetBytes(code);
            _userRepository.Update(user.Id, user);
        }
    }
}
