// MemoryStream
//
// https://www.dotnetperls.com/memorystream

using System;
using System.IO;

class Program
{
  static void Main()
  {
    // Read all bytes in from a file on the disk.
    byte[] file = File.ReadAllBytes("C:\\test.png");
    
    // Create a memory stream from those bytes.
    using (MemoryStream memory = new MemoryStream(file))
    {
      // Use the memory stream from those bytes.
      using (BinaryReader reader = new BinaryReader(memory))
      {
        // Read in each byte from memory.
        for (int i = 0; i < file.Length; i++)
        {
          byte result = reader.ReadByte();
          Console.WriteLine(result);
        }
      }
    }
  }
}
 
