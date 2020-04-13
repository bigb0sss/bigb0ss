using System;
using System.DirectoryServices;
using System.Collections;

namespace find_localAdmin_group
{
    class Program
    {
        static void Main(string[] args)
        {
            DirectoryEntry localMachine = new DirectoryEntry("WinNT://" + Environment.MachineName + ",Computer");
            DirectoryEntry admGroup = localMachine.Children.Find("administrators", "group");
            object members = admGroup.Invoke("members", null);

            Console.WriteLine("[+] Local Administrators: ");
            foreach (object groupMember in (IEnumerable)members)
            {
                DirectoryEntry member = new DirectoryEntry(groupMember);
                Console.WriteLine("\t{0}", member.Name);
            }
            
            Console.Read();
        }
    }
}
