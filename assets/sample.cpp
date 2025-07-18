#include <iostream>
#include <vector>
#include <string>
#include <algorithm>

class Student {
private:
    int id;
    std::string name;
    std::vector<int> grades;
    
public:
    Student(int student_id, const std::string& student_name) 
        : id(student_id), name(student_name) {}
    
    void addGrade(int grade) {
        if (grade >= 0 && grade <= 100) {
            grades.push_back(grade);
        } else {
            std::cout << "Invalid grade: " << grade << std::endl;
        }
    }
    
    double calculateAverage() const {
        if (grades.empty()) {
            return 0.0;
        }
        
        int sum = 0;
        for (int grade : grades) {
            sum += grade;
        }
        
        return static_cast<double>(sum) / grades.size();
    }
    
    int getHighestGrade() const {
        if (grades.empty()) {
            return 0;
        }
        return *std::max_element(grades.begin(), grades.end());
    }
    
    std::string getGradeLevel() const {
        double avg = calculateAverage();
        if (avg >= 90) return "A";
        else if (avg >= 80) return "B";
        else if (avg >= 70) return "C";
        else if (avg >= 60) return "D";
        else return "F";
    }
    
    void displayInfo() const {
        std::cout << "Student ID: " << id << std::endl;
        std::cout << "Name: " << name << std::endl;
        std::cout << "Average: " << calculateAverage() << std::endl;
        std::cout << "Grade level: " << getGradeLevel() << std::endl;
        std::cout << "Highest: " << getHighestGrade() << std::endl;
    }
    
    int getId() const { return id; }
    const std::string& getName() const { return name; }
};

int main() {
    std::cout << "=== Student Grade System ===" << std::endl;
    
    // Create students
    Student alice(1001, "Alice Johnson");
    Student bob(1002, "Bob Smith");
    
    // Add grades for Alice
    alice.addGrade(95);
    alice.addGrade(88);
    alice.addGrade(92);
    
    // Add grades for Bob
    bob.addGrade(78);
    bob.addGrade(82);
    bob.addGrade(75);
    
    // Display student information
    std::cout << "\nStudent Information:" << std::endl;
    std::cout << "-------------------" << std::endl;
    alice.displayInfo();
    
    std::cout << std::endl;
    bob.displayInfo();
    
    std::cout << "\nSystem demo completed!" << std::endl;
    
    return 0;
}
