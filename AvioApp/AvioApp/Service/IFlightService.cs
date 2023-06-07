using AvioApp.Model.DTO;

namespace AvioApp.Service
{
    public interface IFlightService
    {
        string Create(NewFlightDTO newFlight);
        void Delete(string id);
        IEnumerable<FlightAdminPreviewDTO> GetAll();
        IEnumerable<FlightUserPreviewDTO> Search(FlightSearchDTO query);
    }
}
