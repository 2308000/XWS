using AvioApp.Model.DTO;
using AvioApp.Service;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using System.Security.Claims;

namespace AvioApp.Controller
{
    [ApiController]
    [Route("api/[controller]")]
    public class TicketController : ControllerBase
    {
        private readonly ITicketService _ticketService;
        public TicketController(ITicketService ticketService)
        {
            _ticketService = ticketService;
        }

        [HttpPost("{flightId}/{amount}", Name = "Buy")]
        [Authorize(Roles = "USER")]
        public ActionResult<string> Buy(string flightId, int amount)
        {
            _ticketService.Buy(flightId, amount, User.FindFirstValue(ClaimTypes.Email));
            return Ok();
        }

        [HttpGet(Name = "GetAllTickets")]
        [Authorize(Roles = "USER")]
        public ActionResult<IEnumerable<TicketPreviewDTO>> GetAll()
        {
            return Ok(_ticketService.GetAll(User.FindFirstValue(ClaimTypes.Email)));
        }
    }
}
