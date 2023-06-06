namespace AvioApp.SupportClasses.GEH.CustomExceptions
{
    public class NotFoundException : Exception
    {
        public NotFoundException(string message = "Not found!") : base(message)
        { }
    }
}
