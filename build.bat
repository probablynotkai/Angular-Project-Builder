@echo off

rd /s /q dist

go build -o dist/angular-gen.exe

xcopy /s /y templates\* dist\templates\
xcopy /s /y assets\* dist\assets\