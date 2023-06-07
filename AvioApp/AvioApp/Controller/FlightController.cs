using AvioApp.Model.DTO;
using AvioApp.Service;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace AvioApp.Controller
{
    [ApiController]
    [Route("api/[controller]")]
    public class FlightController : ControllerBase
    {
        private readonly IFlightService _flightService;
        public FlightController(IFlightService flightService)
        {
            _flightService = flightService;
        }

        [HttpPost(Name = "Create")]
        [Authorize(Roles = "ADMIN")]
        public ActionResult<long> Create([FromBody] NewFlightDTO newFlight)
        {
            return Ok(_flightService.Create(newFlight));
        }

        [HttpDelete("{id}", Name = "Delete")]
        [Authorize(Roles = "ADMIN")]
        public ActionResult Delete(string id)
        {
            _flightService.Delete(id);
            return Ok();
        }

        [HttpGet(Name = "GetAllFlights")]
        [Authorize(Roles = "ADMIN")]
        public ActionResult<IEnumerable<FlightAdminPreviewDTO>> GetAll()
        {
            return Ok(_flightService.GetAll());
        }

        [HttpPost("search", Name = "Search")]
        [AllowAnonymous]
        public ActionResult<IEnumerable<FlightUserPreviewDTO>> Search([FromBody] FlightSearchDTO query)
        {
            return Ok(_flightService.Search(query));
        }
    }
}
