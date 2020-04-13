using System;
using System.DirectoryServices;

namespace create_new_localuser
{
    class Program
    {
        static void Main(string[] args)
        {
            DirectoryEntry localMachine = new DirectoryEntry("WinNT://" + Environment.MachineName);

            DirectoryEntry newUser = localMachine.Children.Add("test_user", "user");    // Edit "test_user" for username
            newUser.Invoke("SetPassword", new object[] { "Password123" });    // Edit "Password123" for password
            newUser.CommitChanges();
            Console.WriteLine(newUser.Guid.ToString());
            localMachine.Close();
            newUser.Close();
        }
    }
}
