using AvioApp.Model;
using AvioApp.SupportClasses;
using Microsoft.EntityFrameworkCore;

namespace AvioApp.Data
{
    public class XWS_DbContext : DbContext
    {
        public DbSet<User> Users { get; set; }
        public DbSet<Flight> Flights { get; set; }
        public DbSet<Ticket> Tickets { get; set; }
        public XWS_DbContext(DbContextOptions options) : base(options) { }
        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            byte[] salt;
            modelBuilder.Entity<User>().HasKey(u => u.Id);
            modelBuilder.Entity<User>().Property(u => u.Email).IsRequired();
            modelBuilder.Entity<User>().Property(u => u.Password).IsRequired();
            modelBuilder.Entity<User>().Property(u => u.Salt).IsRequired();
            modelBuilder.Entity<User>().Property(u => u.FirstName).IsRequired();
            modelBuilder.Entity<User>().Property(u => u.LastName).IsRequired();
            modelBuilder.Entity<User>().Property(u => u.Role).IsRequired();
            modelBuilder.Entity<User>().Property(u => u.Code).IsRequired(false);
            modelBuilder.Entity<User>(user =>
            {
                user.HasData(
                    new
                    {
                        Id = 1L,
                        Email = "admin@gmail.com",
                        FirstName = "Admin",
                        LastName = "XWS",
                        Password = PasswordHasher.HashPassword("123", out salt),
                        Salt = salt,
                        Role = "ADMIN",
                    }
                );
                user.HasData(
                    new
                    {
                        Id = 2L,
                        Email = "user@gmail.com",
                        FirstName = "User",
                        LastName = "XWS",
                        Password = PasswordHasher.HashPassword("123", out salt),
                        Salt = salt,
                        Role = "USER",
                    }
                );
            });

            modelBuilder.Entity<Flight>().HasKey(f => f.Id);
            modelBuilder.Entity<Flight>().Property(f => f.Date).IsRequired();
            modelBuilder.Entity<Flight>().Property(f => f.Duration).IsRequired();
            modelBuilder.Entity<Flight>().Property(f => f.Start).IsRequired();
            modelBuilder.Entity<Flight>().Property(f => f.Destination).IsRequired();
            modelBuilder.Entity<Flight>().Property(f => f.Price).IsRequired();
            modelBuilder.Entity<Flight>().Property(f => f.Destination).IsRequired();

            modelBuilder.Entity<Ticket>().HasKey(t => t.Id);
            modelBuilder.Entity<Ticket>().HasOne(t => t.Flight).WithMany().IsRequired();
            modelBuilder.Entity<Ticket>().HasOne(t => t.User).WithMany().IsRequired();
        }
    }
}
