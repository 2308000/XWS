using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace AvioApp.Migrations
{
    /// <inheritdoc />
    public partial class AddedRoleToUser : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.AddColumn<string>(
                name: "Role",
                table: "Users",
                type: "nvarchar(max)",
                nullable: false,
                defaultValue: "");

            migrationBuilder.UpdateData(
                table: "Users",
                keyColumn: "Id",
                keyValue: 1L,
                columns: new[] { "Password", "Role", "Salt" },
                values: new object[] { "5A96EE1C23800D0DD409601C62A01E63F70A69F8A8A80536EDA722C74D0265B024F93F4402FBC9FF42569FBD13E785C395E3FB1A5B95A54AE3DB9779D3F8CB59", "ADMIN", new byte[] { 158, 8, 213, 34, 222, 93, 162, 46, 205, 218, 210, 4, 208, 47, 239, 177, 170, 61, 135, 96, 219, 146, 243, 142, 50, 31, 67, 181, 168, 194, 224, 10, 81, 56, 139, 18, 242, 108, 42, 104, 83, 49, 238, 214, 7, 103, 184, 255, 215, 133, 211, 8, 121, 41, 87, 251, 160, 201, 184, 208, 161, 215, 129, 193 } });

            migrationBuilder.UpdateData(
                table: "Users",
                keyColumn: "Id",
                keyValue: 2L,
                columns: new[] { "Password", "Role", "Salt" },
                values: new object[] { "4C0DE2D9AFABD15BCD421E595561D039D99FAD2D23DBF29733702BA43918EE0E85DFE18B029DA7E23B0CE175A944D2733C35F31A086DE9E176C15A6022387023", "USER", new byte[] { 76, 42, 58, 64, 22, 217, 188, 93, 167, 250, 18, 14, 236, 140, 5, 128, 151, 139, 129, 245, 198, 190, 133, 141, 191, 233, 101, 158, 168, 175, 125, 2, 132, 143, 148, 21, 198, 98, 221, 121, 26, 3, 66, 53, 44, 169, 75, 59, 202, 6, 243, 94, 247, 238, 213, 206, 109, 190, 148, 207, 154, 30, 221, 117 } });
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropColumn(
                name: "Role",
                table: "Users");

            migrationBuilder.UpdateData(
                table: "Users",
                keyColumn: "Id",
                keyValue: 1L,
                columns: new[] { "Password", "Salt" },
                values: new object[] { "37E1F4057F63E21961E3A11249F17992F6898BABCDE143C9D9E229796C46C4A77D53D250EF9808D6F532198562D8998E47945D0DDB2237C14A678541C6B95A3A", new byte[] { 125, 194, 71, 144, 23, 133, 128, 21, 182, 118, 126, 169, 162, 59, 40, 227, 242, 229, 104, 52, 128, 140, 42, 34, 207, 253, 121, 80, 219, 22, 67, 72, 185, 146, 159, 7, 67, 91, 39, 160, 122, 115, 124, 0, 167, 65, 140, 188, 15, 133, 12, 22, 145, 232, 110, 227, 216, 183, 30, 162, 82, 96, 243, 155 } });

            migrationBuilder.UpdateData(
                table: "Users",
                keyColumn: "Id",
                keyValue: 2L,
                columns: new[] { "Password", "Salt" },
                values: new object[] { "8CE7D420515B3724CB837FE4EAD325C5FFF1A7837874C2CAE08B95BD577B5338F82636F965E025BF270BDB561FBB84B96A314F4CFE4BE4640B9D85743B7DADC9", new byte[] { 190, 61, 250, 132, 1, 2, 104, 106, 190, 1, 174, 221, 221, 194, 117, 218, 176, 80, 205, 79, 78, 171, 177, 128, 220, 126, 236, 84, 109, 134, 91, 253, 15, 51, 236, 135, 164, 158, 163, 214, 88, 11, 238, 174, 16, 70, 248, 54, 185, 129, 80, 186, 240, 99, 97, 147, 193, 98, 208, 171, 75, 157, 162, 159 } });
        }
    }
}
