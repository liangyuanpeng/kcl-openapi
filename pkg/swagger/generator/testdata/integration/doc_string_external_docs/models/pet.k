"""
This file was generated by the KCL auto-gen tool. DO NOT EDIT.
Editing this file might prove futile when you re-run the KCL auto-gen generate command.
"""


schema Pet:
    r"""
    pet

    Attributes
    ----------
    name : str, default is Undefined, optional
        the name of the pet. must be a RFC 1123 name. See Also: https://datatracker.ietf.org/doc/html/rfc1123
    age : int, default is Undefined, optional
        age
    """


    name?: str

    age?: int


