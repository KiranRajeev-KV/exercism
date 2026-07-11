// this is package twofer
package twofer

// ShareWith is a function
func ShareWith(name string) string {
    if(name == "") {
        return "One for you, one for me."
    }
    return "One for " + name + ", one for me."
    
}
