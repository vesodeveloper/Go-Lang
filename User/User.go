package main
import "fmt"
type Teacher struct{
    ID int
    Name string
    Age int
    Departament string
}
type Student struct{
    Name string
    Age int
    Spec string
    FNum string
}
func IntroName(t1 Teacher){
     fmt.Println("Name is:")
    fmt.Println(t1.Name)
}
func IntroAge(t1 Teacher){
    fmt.Println("Age is:")
    fmt.Println(t1.Age)
}
func IntroDep(t1 Teacher){
     fmt.Println("Departament is:")
    fmt.Println(t1.Departament)
}
func main(){
//Teacher Instances
  teacher1:=Teacher{ 
      ID:01,
      Name:"Stoyan Cheresharov",
    Age:  35,
     Departament: "Software Technologies",
      
} 
  teacher2:=Teacher{
    ID:  02,
     Name: "Svetoslav Enkov",
    Age:  34,
     Departament: "Computer Informatics",}
  teacher3:=Teacher{ 
     ID: 03,
     Name: "Kiril Ivanov",
     Age: 45,
    Departament:  "Computer Informatics",}
  //Student Instances
  student1:=Student{
      Name:"Kiril Kirilov",
      Age:18,
      Spec:"Software Engineering",
      FNum:"1801322345",
}
  student2:=Student{
      Name:"Peter Petrov",
      Age:18,
      Spec:"Software Design",
      FNum:"1801322378",
}
student3:=Student{
      Name:"Stefan Stefanov",
      Age:18,
      Spec:"Informatics",
      FNum:"1801322354",
}
//Make some changes
student1.Spec="Information Technologies"
student2.Age=19
student3.Name="Mihail Hristov"
if student2.Age==19{
    fmt.Println("Age is Correct")
}else{
    fmt.Println("Age is not Correct!")
}
//Executable Intro function
IntroAge(teacher1)
IntroName(teacher2)
IntroDep(teacher3)
}
