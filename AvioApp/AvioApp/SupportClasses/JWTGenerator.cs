using AvioApp.Model;
using Microsoft.IdentityModel.Tokens;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;

namespace AvioApp.SupportClasses
{
    public class JWTGenerator : IJWTGenerator
    {
        private readonly IConfiguration _config;
        readonly JwtSecurityTokenHandler _jwtHandler = new();
        public JWTGenerator(IConfiguration config)
        {
            _config = config;
        }

        public string GenerateToken(User user, bool isLoginPermanent = false)
        {
            var securityKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(_config["Jwt:Key"]));
            var credentials = new SigningCredentials(securityKey, SecurityAlgorithms.HmacSha256);
            var claims = new[]
            {
                new Claim(ClaimTypes.Email, user.Email),
                new Claim(ClaimTypes.Role, user.Role)
            };

            SecurityTokenDescriptor tokenDescriptor = new SecurityTokenDescriptor
            {
                Subject = new ClaimsIdentity(claims),
                Expires = isLoginPermanent ? DateTime.UtcNow.AddYears(10) : DateTime.UtcNow.AddMinutes(30),
                SigningCredentials = credentials,
                Issuer = _config["Jwt:Issuer"],
                Audience = _config["Jwt:Audience"]
            };

            SecurityToken token = _jwtHandler.CreateToken(tokenDescriptor);
            return _jwtHandler.WriteToken(token);
        }
    }
}
