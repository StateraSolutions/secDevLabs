const User = require('./models/user');
var sanitize = require('mongo-sanitize');

const register = async (user) => {

    try { 
        const { name, email, password } = user;

        const existUser = await User.findOne({ email: sanitize(email) });

        if(existUser) { return null }

        const newUser = new User({
            name: name,
            email: email,
            password: password
        });

        await newUser.save();

        return newUser;
    }

    catch(error) { throw error; }
    
}

const login = async (credentials) => {

    try {
        const { email, password } = credentials;

        const existsUser = await User.find({$and: [ { email: sanitize(email)}, { password: sanitize(password)} ]});

        if(!existsUser) { return null;}

        const returnUser = existsUser.map((user) => {
            return user.email
        })


        return returnUser;
    }

    catch(error) { throw error; }
    

}

module.exports = {
    register,
    login
};