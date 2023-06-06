using AvioApp.Model;

namespace AvioApp.SupportClasses
{
    public interface IJWTGenerator
    {
        string GenerateToken(User user, bool isLoginPermanent = false);
    }
}
