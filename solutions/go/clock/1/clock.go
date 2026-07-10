package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
    hour int
    min int
}

func New(h, m int) Clock {
    c := Clock{
        hour: m/60 +h,  
        min: m%60,
    }

    if(c.min >=60) {
        c.hour = c.min/60
        c.min = c.min%60
    } else if (c.min <=-1) {
        c.hour = c.hour + c.min/60 -1
        c.min = 60 + c.min%60
    }

        if(c.hour >=24) {
        c.hour = c.hour%24
    } else if (c.hour <=-1) {
        c.hour = 24 + c.hour%24
    }
    
    return c
}

func (c Clock) Add(m int) Clock {
    c.min += m%60
    c.hour += m/60
    
    if(c.min >=60) {
        c.hour += c.min/60
        c.min = c.min%60
    } else if (c.min <=-1) {
        c.hour = c.hour + c.min/60 -1
        c.min = 60 + c.min%60
    }

    if(c.hour >=24) {
        c.hour = c.hour%24
    } else if (c.hour <=-1) {
        c.hour = 24 + c.hour%24
    }
    return c
}

func (c Clock) Subtract(m int) Clock {
    c.min -= m%60 // 20
    c.hour -= m/60 // -48

    if(c.min >=60) {
        c.hour += c.min/60 
        c.min = c.min%60
    } else if (c.min <=-1) {
        c.hour = c.hour + c.min/60 -1 //-49.5
        c.min = 60 + c.min%60 // 30
    }

    if(c.hour >=24) {
        c.hour = c.hour%24
    } else if (c.hour <=-1) {
        c.hour = 24 + c.hour%24
    }

    if(c.hour >=24) {
        c.hour = c.hour%24
    } else if (c.hour <=-1) {
        c.hour = 24 + c.hour%24
    }
    return c
}

func (c Clock) String() string {
    return fmt.Sprintf("%02d:%02d",c.hour,c.min)
}
