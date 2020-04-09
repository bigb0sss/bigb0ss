using System;
using System.Net;

namespace Hostname
{
    class Program
    {
        static void Main(string[] args)
        {
            string hostName = Dns.GetHostName();
            Console.Out.WriteLine("Hostname: {0}", hostName);
            Console.Read();
        }
    }
}
