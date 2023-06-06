using AvioApp.SupportClasses.GEH.CustomExceptions;
using System.Net;

namespace AvioApp.SupportClasses.GEH
{
    public class ExceptionDictionary
    {
        private static Dictionary<Type, HttpStatusCode> exceptionStatusCodes = new Dictionary<Type, HttpStatusCode>
        {
            {typeof(Exception), HttpStatusCode.InternalServerError},
            {typeof(NotFoundException), HttpStatusCode.NotFound},
            {typeof(BadCredentialsException), HttpStatusCode.BadRequest},
            {typeof(DuplicateItemException), HttpStatusCode.BadRequest},
            {typeof(InsufficientResoucesException), HttpStatusCode.BadRequest}
        };

        public static HttpStatusCode GetExceptionStatusCode(Exception ex)
        {
            bool exceptionFound = exceptionStatusCodes.TryGetValue(ex.GetType(), out var statusCode);
            return exceptionFound ? statusCode : HttpStatusCode.InternalServerError;
        }
    }
}
