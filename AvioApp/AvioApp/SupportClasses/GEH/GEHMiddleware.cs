using System.Net;
using System.Text.Json;

namespace AvioApp.SupportClasses.GEH
{
    public class GEHMiddleware
    {
        private readonly RequestDelegate _next;

        public GEHMiddleware(RequestDelegate next)
        {
            _next = next;
        }

        public async Task Invoke(HttpContext context)
        {
            try
            {
                await _next(context);
            }
            catch (Exception ex)
            {
                HttpStatusCode httpStatusCode = ExceptionDictionary.GetExceptionStatusCode(ex);

                context.Response.StatusCode = (int)httpStatusCode;
                var exceptionResult = JsonSerializer.Serialize(new { message = ex.Message });
                context.Response.ContentType = "application/json";
                await context.Response.WriteAsync(exceptionResult);

            }
        }
    }
}
