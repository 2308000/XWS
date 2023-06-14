using AvioApp.Model;
using AvioApp.Repository;
using AvioApp.SupportClasses;

namespace AvioApp.Service
{
    public class InitService : IInitService
    {
        private readonly IFlightRepository _flightRepository;
        private readonly ITicketRepository _ticketRepository;
        private readonly IUserRepository _userRepository;

        public InitService(IFlightRepository flightRepository, ITicketRepository ticketRepository, IUserRepository userRepository)
        {
            _flightRepository = flightRepository;
            _ticketRepository = ticketRepository;
            _userRepository = userRepository;
        }

        public void Init()
        {
            var random = new Random();
            var users = _userRepository.GetAll();
            foreach (var user in users)
            {
                _userRepository.Delete(user.Id);
            }

            {
                byte[] salt;
                _userRepository.Create(new User
                {
                    Email = "nikola@gmail.com",
                    Password = PasswordHasher.HashPassword("123", out salt),
                    Salt = salt,
                    Role = "USER",
                    FirstName = "Nikola",
                    LastName = "Vukic"
                });
                _userRepository.Create(new User
                {
                    Email = "dane@gmail.com",
                    Password = PasswordHasher.HashPassword("123", out salt),
                    Salt = salt,
                    Role = "USER",
                    FirstName = "Dane",
                    LastName = "Milisic"
                });
                _userRepository.Create(new User
                {
                    Email = "admin@gmail.com",
                    Password = PasswordHasher.HashPassword("123", out salt),
                    Salt = salt,
                    Role = "ADMIN",
                    FirstName = "ADMIN",
                    LastName = "XWS"
                });
            }



            var flights = _flightRepository.GetAll();
            foreach (var flight in flights)
            {
                _flightRepository.Delete(flight.Id);
            }

            {
                var cities = new List<string>() { "Beograd", "Podgorica", "Maribor", "Budimpesta" };
                for (int i = 0; i < 12; i++)
                {
                    var num1 = random.Next(cities.Count);
                    var num2 = num1;
                    while (num1 == num2)
                    {
                        num2 = random.Next(cities.Count);
                    }
                    _flightRepository.Create(new Flight
                    {
                        Date = DateTime.Today.Date.AddDays(random.Next(1, 3)).AddHours(random.Next(8, 20)),
                        Duration = random.Next(30, 46),
                        Start = cities[num1],
                        Destination = cities[num2],
                        Price = random.Next(80, 110),
                        Tickets = random.Next(3, 10)
                    });
                }
            }



            var tickets = _ticketRepository.GetAll();
            foreach (var ticket in tickets)
            {
                _ticketRepository.Delete(ticket.Id);
            }

            {
                flights = _flightRepository.GetAll();
                var user = _userRepository.GetAll().First();
                foreach (var flight in flights)
                {
                    if (random.Next(2) == 1)
                    {
                        var cap = random.Next(4);
                        for (int i = 0; i < cap; i++)
                        {
                            _ticketRepository.Create(new Ticket
                            {
                                FlightId = flight.Id,
                                UserId = user.Id,
                            });
                        }
                    }
                }
            }
        }
    }
}
