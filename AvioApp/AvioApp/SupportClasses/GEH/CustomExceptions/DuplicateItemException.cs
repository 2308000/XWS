namespace AvioApp.SupportClasses.GEH.CustomExceptions
{
    public class DuplicateItemException : Exception
    {
        public DuplicateItemException(string message = "This item already exists in database.") : base(message) { }
    }
}
