using AvioApp.Model.DTO;
using AvioApp.Service;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using System.Security.Claims;

namespace AvioApp.Controllers
{
    [ApiController]
    [Route("api/[controller]")]
    public class UserController : ControllerBase
    {
        private readonly IUserService _userService;
        public UserController(IUserService userService)
        {
            _userService = userService;
        }

        [HttpPost("login", Name = "Login")]
        [AllowAnonymous]
        public ActionResult<string> Login([FromBody] CredentialsDTO credentials)
        {
            return Ok(_userService.Authenticate(credentials));
        }

        [HttpPost("registration", Name = "Registration")]
        [AllowAnonymous]
        public ActionResult<long> Registration([FromBody] NewUserDTO newUser)
        {
            return Ok(_userService.Register(newUser));
        }

        [HttpPatch("code/{code}", Name = "UpdateCode")]
        [Authorize(Roles = "USER")]
        public ActionResult UpdateCode(string code)
        {
            _userService.UpdateCode(code, User.FindFirstValue(ClaimTypes.Email));
            return Ok();
        }

    }
}