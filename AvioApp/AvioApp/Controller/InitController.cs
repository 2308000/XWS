using AvioApp.Service;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace AvioApp.Controller
{
    [ApiController]
    [Route("api/[controller]")]
    public class InitController : ControllerBase
    {
        private readonly IInitService _initService;
        public InitController(IInitService initService)
        {
            _initService = initService;
        }

        [HttpPost(Name = "Init")]
        [AllowAnonymous]
        public ActionResult Create()
        {
            _initService.Init();
            return Ok();
        }
    }
}
