usudo
=====

A ``sudo`` for Windows, designed to handle both command-line and windowed
applications well.

What's wrong with the others?
*****************************

Nothing in particular. However, what I wanted was:

- Something that didn't depend on PowerShell or VBScript. It should be a
  self-contained binary.

- Something that worked well with command-line programs, e.g. pausing after
  the application finishes so you can actually see any errors/messages.

Downloads
*********

All binaries (for X86 *and* X64) are available at
`files.fm <https://files.fm/u/2vdbn2xs>`_.

How to use
**********

Toss up all 3 executables up somewhere in your PATH. The two important ones are
``usudo.exe`` and ``usudo-w.exe``.

``usudo.exe`` will *always* open a command prompt, and it will ``pause`` once
the application finishes running. ``usudo-w.exe`` will NOT open a command prompt
unless necessary (e.g. running a command-line-only application), and there's
kind of pause after it finishes.

**TL;DR:** Use ``usudo.exe`` for command-line applications and ``usudo-w.exe``
for windowed applications.

Example usage::

  usudo choco install googlechrome
  usudo-w notepad some-important-file.txt
