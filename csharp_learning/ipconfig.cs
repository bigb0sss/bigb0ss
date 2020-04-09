using System;
using System.Net;

namespace ipconfig
{
    class Program
    {
        static void Main(string[] args)
        {
            // Getting Hostname
            string hostName = Dns.GetHostName();    
            Console.Out.WriteLine("Hostname: {0}", hostName);

            // Getting IP Addres
            string ipInfo = Dns.GetHostByName(hostName).AddressList[0].ToString();
            Console.Out.WriteLine("IP Address: {0}", ipInfo);
            Console.Read();
        }
    }
}
