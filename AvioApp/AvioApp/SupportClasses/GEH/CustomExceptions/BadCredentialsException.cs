namespace AvioApp.SupportClasses.GEH.CustomExceptions
{
    public class BadCredentialsException : Exception
    {
        public BadCredentialsException(string message = "Credentials do not match!") : base(message)
        { }
    }
}
