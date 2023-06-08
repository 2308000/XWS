#Build Stage
FROM mcr.microsoft.com/dotnet/sdk:6.0-focal AS build
RUN dotnet dev-certs https
WORKDIR /source
COPY . .

#COPY xwscert.pfx /app/xwscert.pfx
#RUN apt-get update && apt-get install -y openssl
#RUN openssl pkcs12 -in /app/xwscert.pfx -nokeys -out /app/xwscert.crt -nodes -password pass:123
#RUN cp /app/xwscert.crt /usr/local/share/ca-certificates/
#RUN update-ca-certificates

RUN dotnet restore "./AvioApp/AvioApp/AvioApp.csproj" --disable-parallel
RUN dotnet publish "./AvioApp/AvioApp/AvioApp.csproj" -c release -o /app --no-restore

#Serve Stage
FROM mcr.microsoft.com/dotnet/aspnet:6.0-focal
WORKDIR /app
COPY --from=build /app ./

COPY --from=build /root/.dotnet/corefx/cryptography/x509stores/my/* /root/.dotnet/corefx/cryptography/x509stores/my/

EXPOSE 5000

ENTRYPOINT ["dotnet", "AvioApp.dll"]