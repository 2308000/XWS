using AvioApp.Model.DTO;

namespace AvioApp.Service
{
    public interface IFlightService
    {
        long Create(NewFlightDTO newFlight);
        void Delete(long id);
        IEnumerable<FlightAdminPreviewDTO> GetAll();
        IEnumerable<FlightUserPreviewDTO> Search(FlightSearchDTO query);
    }
}
