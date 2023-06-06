using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

#pragma warning disable CA1814 // Prefer jagged arrays over multidimensional

namespace AvioApp.Migrations
{
    /// <inheritdoc />
    public partial class AddedUserModel : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "Users",
                columns: table => new
                {
                    Id = table.Column<long>(type: "bigint", nullable: false)
                        .Annotation("SqlServer:Identity", "1, 1"),
                    Email = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Password = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Salt = table.Column<byte[]>(type: "varbinary(max)", nullable: false),
                    FirstName = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    LastName = table.Column<string>(type: "nvarchar(max)", nullable: false),
                    Code = table.Column<byte[]>(type: "varbinary(max)", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Users", x => x.Id);
                });

            migrationBuilder.InsertData(
                table: "Users",
                columns: new[] { "Id", "Code", "Email", "FirstName", "LastName", "Password", "Salt" },
                values: new object[,]
                {
                    { 1L, null, "admin@gmail.com", "Admin", "XWS", "37E1F4057F63E21961E3A11249F17992F6898BABCDE143C9D9E229796C46C4A77D53D250EF9808D6F532198562D8998E47945D0DDB2237C14A678541C6B95A3A", new byte[] { 125, 194, 71, 144, 23, 133, 128, 21, 182, 118, 126, 169, 162, 59, 40, 227, 242, 229, 104, 52, 128, 140, 42, 34, 207, 253, 121, 80, 219, 22, 67, 72, 185, 146, 159, 7, 67, 91, 39, 160, 122, 115, 124, 0, 167, 65, 140, 188, 15, 133, 12, 22, 145, 232, 110, 227, 216, 183, 30, 162, 82, 96, 243, 155 } },
                    { 2L, null, "user@gmail.com", "User", "XWS", "8CE7D420515B3724CB837FE4EAD325C5FFF1A7837874C2CAE08B95BD577B5338F82636F965E025BF270BDB561FBB84B96A314F4CFE4BE4640B9D85743B7DADC9", new byte[] { 190, 61, 250, 132, 1, 2, 104, 106, 190, 1, 174, 221, 221, 194, 117, 218, 176, 80, 205, 79, 78, 171, 177, 128, 220, 126, 236, 84, 109, 134, 91, 253, 15, 51, 236, 135, 164, 158, 163, 214, 88, 11, 238, 174, 16, 70, 248, 54, 185, 129, 80, 186, 240, 99, 97, 147, 193, 98, 208, 171, 75, 157, 162, 159 } }
                });
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "Users");
        }
    }
}
