namespace AvioApp.SupportClasses.GEH.CustomExceptions
{
    public class InsufficientResoucesException : Exception
    {
        public InsufficientResoucesException(string message = "Not enough resources!") : base(message)
        { }
    }
}
