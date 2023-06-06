using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace AvioApp.Migrations
{
    /// <inheritdoc />
    public partial class AddedFlightModel : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "Flights",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("SqlServer:Identity", "1, 1"),
                    Date = table.Column<DateTime>(type: "datetime2", nullable: false),
                    Duration = table.Column<int>(type: "int", nullable: false),
                    Start = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Destination = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Price = table.Column<float>(type: "real", nullable: false),
                    Tickets = table.Column<int>(type: "int", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Flights", x => x.Id);
                });

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

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "Flights");

            migrationBuilder.UpdateData(
                table: "Users",
                keyColumn: "Id",
                keyValue: 1L,
                columns: new[] { "Password", "Salt" },
                values: new object[] { "5A96EE1C23800D0DD409601C62A01E63F70A69F8A8A80536EDA722C74D0265B024F93F4402FBC9FF42569FBD13E785C395E3FB1A5B95A54AE3DB9779D3F8CB59", new byte[] { 158, 8, 213, 34, 222, 93, 162, 46, 205, 218, 210, 4, 208, 47, 239, 177, 170, 61, 135, 96, 219, 146, 243, 142, 50, 31, 67, 181, 168, 194, 224, 10, 81, 56, 139, 18, 242, 108, 42, 104, 83, 49, 238, 214, 7, 103, 184, 255, 215, 133, 211, 8, 121, 41, 87, 251, 160, 201, 184, 208, 161, 215, 129, 193 } });

            migrationBuilder.UpdateData(
                table: "Users",
                keyColumn: "Id",
                keyValue: 2L,
                columns: new[] { "Password", "Salt" },
                values: new object[] { "4C0DE2D9AFABD15BCD421E595561D039D99FAD2D23DBF29733702BA43918EE0E85DFE18B029DA7E23B0CE175A944D2733C35F31A086DE9E176C15A6022387023", new byte[] { 76, 42, 58, 64, 22, 217, 188, 93, 167, 250, 18, 14, 236, 140, 5, 128, 151, 139, 129, 245, 198, 190, 133, 141, 191, 233, 101, 158, 168, 175, 125, 2, 132, 143, 148, 21, 198, 98, 221, 121, 26, 3, 66, 53, 44, 169, 75, 59, 202, 6, 243, 94, 247, 238, 213, 206, 109, 190, 148, 207, 154, 30, 221, 117 } });
        }
    }
}
