using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace AvioApp.Migrations
{
    /// <inheritdoc />
    public partial class AddedTicketModel : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "Tickets",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("SqlServer:Identity", "1, 1"),
                    UserId = table.Column<long>(type: "bigint", nullable: false),
                    FlightId = table.Column<long>(type: "bigint", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Tickets", x => x.Id);
                    table.ForeignKey(
                        name: "FK_Tickets_Flights_FlightId",
                        column: x => x.FlightId,
                        principalTable: "Flights",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                    table.ForeignKey(
                        name: "FK_Tickets_Users_UserId",
                        column: x => x.UserId,
                        principalTable: "Users",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.UpdateData(
                table: "Users",
                keyColumn: "Id",
                keyValue: 1L,
                columns: new[] { "Password", "Salt" },
                values: new object[] { "8557DFC3F4047B55F861C46AE7F75E23406C9CE5C6002CEC56E903DD139EB25A78F7A2E5BF7A2CBD5199EF40A9A9B67C1934D2CDF2B57351C6D611DEDDCB6F8E", new byte[] { 180, 90, 161, 11, 251, 122, 18, 70, 185, 129, 195, 213, 135, 250, 82, 3, 12, 150, 16, 142, 56, 219, 203, 206, 137, 135, 60, 199, 28, 141, 232, 137, 77, 87, 23, 150, 146, 252, 13, 236, 137, 217, 72, 91, 43, 205, 115, 19, 119, 138, 249, 13, 89, 124, 186, 65, 156, 20, 86, 107, 102, 88, 104, 73 } });

            migrationBuilder.UpdateData(
                table: "Users",
                keyColumn: "Id",
                keyValue: 2L,
                columns: new[] { "Password", "Salt" },
                values: new object[] { "F737363D9240FF0CEA1B1846E487572594EEC89339EACCDDF2072AB8CB839758294C03EB94B6A3957A441F7FFDDF18F9618A92DCBA8E3092E5E5287E3B4A0FC2", new byte[] { 76, 147, 52, 244, 64, 34, 87, 2, 68, 27, 68, 51, 206, 156, 71, 138, 245, 103, 248, 225, 16, 34, 203, 95, 203, 163, 50, 219, 58, 40, 159, 144, 73, 18, 223, 181, 203, 50, 19, 44, 172, 17, 232, 113, 94, 194, 174, 117, 35, 148, 132, 5, 73, 53, 245, 183, 225, 201, 7, 94, 164, 236, 78, 230 } });

            migrationBuilder.CreateIndex(
                name: "IX_Tickets_FlightId",
                table: "Tickets",
                column: "FlightId");

            migrationBuilder.CreateIndex(
                name: "IX_Tickets_UserId",
                table: "Tickets",
                column: "UserId");
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "Tickets");

            migrationBuilder.UpdateData(
                table: "Users",
                keyColumn: "Id",
                keyValue: 1L,
                columns: new[] { "Password", "Salt" },
                values: new object[] { "EAA5FB65866C33255AF997A6486CF58894C3BDF23D9D38B08C7952516C997A70667C7C6CBCE2C13A5D4FC53BA9FC2BE13862929E4A45090369D2684ABDC3EDD9", new byte[] { 201, 39, 223, 251, 166, 117, 20, 173, 223, 10, 252, 5, 32, 157, 254, 87, 118, 55, 246, 24, 76, 71, 18, 31, 131, 3, 229, 145, 242, 68, 52, 78, 188, 46, 248, 142, 227, 131, 235, 37, 102, 140, 10, 240, 115, 172, 56, 20, 131, 129, 101, 188, 109, 80, 130, 199, 165, 9, 32, 111, 133, 64, 195, 198 } });

            migrationBuilder.UpdateData(
                table: "Users",
                keyColumn: "Id",
                keyValue: 2L,
                columns: new[] { "Password", "Salt" },
                values: new object[] { "264B5AB5C5B264E322292FD847D315DDD5BF5215249229430A2D93B9F7D7F50F61A3758CF679B5C71F1878A866EBA0F3263904B5B31FF6B0215AA16EC5230ED8", new byte[] { 120, 49, 170, 250, 214, 193, 210, 103, 225, 74, 127, 33, 115, 190, 29, 172, 80, 143, 72, 142, 185, 82, 175, 98, 162, 249, 91, 191, 192, 212, 82, 111, 77, 56, 48, 242, 30, 18, 2, 3, 238, 174, 171, 61, 141, 122, 135, 71, 90, 156, 34, 2, 153, 128, 23, 145, 82, 167, 72, 219, 46, 182, 146, 208 } });
        }
    }
}
